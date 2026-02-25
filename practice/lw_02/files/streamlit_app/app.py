import streamlit as st
import pandas as pd
import sqlite3
import plotly.express as px
import numpy as np

st.set_page_config(layout="wide")

# -----------------------
# Загрузка данных
# -----------------------

conn = sqlite3.connect("data/clients.db")
df = pd.read_sql("SELECT * FROM clients", conn)

df["debt_to_income"] = df["debt"] / df["income"]
# -----------------------
# Sidebar фильтры
# -----------------------

st.sidebar.header("Фильтры")

age_range = st.sidebar.slider(
    "Возраст",
    int(df.age.min()),
    int(df.age.max()),
    (int(df.age.min()), int(df.age.max()))
)

score_range = st.sidebar.slider(
    "Кредитный рейтинг",
    int(df.credit_score.min()),
    int(df.credit_score.max()),
    (int(df.credit_score.min()), int(df.credit_score.max()))
)

show_delinquent = st.sidebar.selectbox(
    "Просрочка",
    ["Все", "Только с просрочкой", "Без просрочки"]
)

filtered_df = df[
    (df.age.between(age_range[0], age_range[1])) &
    (df.credit_score.between(score_range[0], score_range[1]))
]

if show_delinquent == "Только с просрочкой":
    filtered_df = filtered_df[filtered_df.delinquent == 1]
elif show_delinquent == "Без просрочки":
    filtered_df = filtered_df[filtered_df.delinquent == 0]
# -----------------------
# KPI блок
# -----------------------

col1, col2, col3, col4 = st.columns(4)

col1.metric("Всего клиентов", len(filtered_df))
col2.metric("Средний рейтинг", round(filtered_df.credit_score.mean(), 1))
col3.metric("Средний доход", f"{round(filtered_df.income.mean(), 0)}")
col4.metric("Доля просрочек", f"{round(filtered_df.delinquent.mean()*100, 1)} %")

st.divider()
# -----------------------
# Графики
# -----------------------

col1, col2 = st.columns(2)

with col1:
    fig_score = px.histogram(
        filtered_df,
        x="credit_score",
        nbins=30,
        title="Распределение кредитного рейтинга",
        color="delinquent",
        color_discrete_map={0: "blue", 1: "red"}
    )
    st.plotly_chart(fig_score, use_container_width=True)

with col2:
    fig_income = px.histogram(
        filtered_df,
        x="income",
        nbins=30,
        title="Распределение дохода"
    )
    st.plotly_chart(fig_income, use_container_width=True)

# -----------------------
# Scatter интерактивный
# -----------------------

fig_scatter = px.scatter(
    filtered_df,
    x="income",
    y="debt",
    color="delinquent",
    size="debt_to_income",
    hover_data=["age", "credit_score"],
    title="Debt vs Income (размер = Debt/Income)",
    color_discrete_map={0: "blue", 1: "red"}
)

st.plotly_chart(fig_scatter, use_container_width=True)
# -----------------------
# Корреляция
# -----------------------

st.subheader("Корреляционная матрица")

corr = filtered_df[["age", "income", "debt", "credit_score", "debt_to_income"]].corr()

fig_corr = px.imshow(
    corr,
    text_auto=True,
    title="Correlation Matrix"
)

st.plotly_chart(fig_corr, use_container_width=True)

# -----------------------
# Сегментация риска
# -----------------------
conditions = [
    (filtered_df.credit_score >= 700),
    (filtered_df.credit_score.between(600, 699)),
    (filtered_df.credit_score < 600)
]

choices = ["Low Risk", "Medium Risk", "High Risk"]

filtered_df["risk_segment"] = np.select(
    conditions,
    choices,
    default="Unknown"
)

segment_count = filtered_df["risk_segment"].value_counts().reset_index()
segment_count.columns = ["Segment", "Count"]
fig_segment = px.pie(
    segment_count,
    names="Segment",
    values="Count",
    title="Распределение по сегментам риска"
)

st.plotly_chart(fig_segment, use_container_width=True)