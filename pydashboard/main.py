import streamlit as st
from pydantic import BaseModel, TypeAdapter
from typing import Final
import json
import pandas as pd
import numpy as np
import requests
from requests.exceptions import ConnectionError, HTTPError

URL: Final[str] = "http://localhost:8601"


class TodoItem(BaseModel):
    id: int
    description: str
    done: bool


TodoList = TypeAdapter(list[TodoItem])


def sort_data(data: list) -> list | None:
    response = requests.post(f"{URL}/sort", data=json.dumps({"unsorted_data": data}))
    try:
        response.raise_for_status()
    except (ConnectionError, HTTPError) as e:
        st.markdown(f"Error with sorting: {e}")
        return None

    data = response.json()["sorted_data"]
    return data


def get_todo() -> list[TodoItem]:
    response = requests.get(f"{URL}/todo")
    try:
        response.raise_for_status()
    except (ConnectionError, HTTPError) as e:
        print(f"Error when requesting ToDo items: {e}")
        return []

    todo_items = response.json()
    todo_items = TodoList.validate_python(todo_items)

    return todo_items


def insert_new_item():
    new_description = st.text_input("Description of task")
    if new_description:
        response = requests.post(
            f"{URL}/todo", data=json.dumps({"description": new_description})
        )
        try:
            response.raise_for_status()
        except (ConnectionError, HTTPError) as e:
            print(f"Error when inserting todo item: {e}")
        else:
            st.text("Item submitted")


def toggle_item(item: TodoItem):
    response = requests.put(
        f"{URL}/todo", data=json.dumps({"id": item.id, "done": not item.done})
    )
    try:
        response.raise_for_status()
    except (ConnectionError, HTTPError) as e:
        print(f"Error on update: {e}")


def display_todo_items(items: list[TodoItem]):
    for i, item in enumerate(items):
        description_col, done_col = st.columns([3, 1])
        with description_col:
            st.text(item.description)
        with done_col:
            st.checkbox(
                "",
                value=item.done,
                key=i,
                on_change=toggle_item,
                args=(item,),
                label_visibility="hidden",
            )


def run_app():
    st.title("Some sorted data")

    insert_new_item()

    todo_items = get_todo()
    display_todo_items(todo_items)

    data = list(np.random.randn(25))
    data = sort_data(data)
    if data is None:
        st.markdown("No data")
    else:
        data = pd.Series(data)
        st.table(data)


if __name__ == "__main__":
    run_app()
