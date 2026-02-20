import streamlit as st
import json
import pandas as pd
import numpy as np
import requests


def run_app():
    st.title("Some sorted data")

    data = list(np.random.randn(25))
    response = requests.post(
        "http://localhost:8601/sort", data=json.dumps({"unsorted_data": data})
    )
    try:
        response.raise_for_status()
    except Exception as e:
        st.markdown(f"Error with sorting: {e}")
        return

    data = response.json()["sorted_data"]
    if data is None:
        st.markdown("No data")
    else:
        data = pd.Series(data)
        st.table(data)


if __name__ == "__main__":
    run_app()
