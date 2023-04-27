import grpc
import streamlit as st
import user_pb2
import user_pb2_grpc

st.title("APIテスト画面 ユーザー")

with st.form(key="user"):
    name: str = st.text_input("ユーザー名", max_chars=12)
    submit_button = st.form_submit_button(label="送信")

if submit_button:
    with grpc.insecure_channel("127.0.0.1:8000") as channel:
        stub = user_pb2_grpc.UserServiceStub(channel)
        response = stub.Create(user_pb2.CreateReqeust(user=user_pb2.NewUser(name=name)))
    st.write(response)
