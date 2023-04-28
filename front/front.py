import os

import grpc
import room_pb2
import room_pb2_grpc
import streamlit as st
import user_pb2
import user_pb2_grpc


def run():
    page = st.sidebar.selectbox("Choose your page", ["bookings", "rooms", "users"])
    st.title("APIテスト画面 ユーザー")
    handlers = {"rooms": show_rooms_page, "users": show_users_page}
    handlers[page]()


def show_rooms_page() -> None:
    room_service = os.environ["ROOM"]

    with st.form(key="room"):
        name: str = st.text_input("会議室名", max_chars=12)
        capacity: int = st.number_input("定員", step=1)
        submit_button = st.form_submit_button(label="送信")

    if submit_button:
        with grpc.insecure_channel(room_service) as channel:
            stub = room_pb2_grpc.RoomServiceStub(channel)
            response = stub.Create(
                room_pb2.CreateReqeust(
                    room=room_pb2.NewRoom(name=name, capacity=capacity)
                )
            )
        st.write(response)


def show_users_page() -> None:
    user_service = os.environ["USER"]

    with st.form(key="user"):
        name: str = st.text_input("ユーザー名", max_chars=12)
        submit_button = st.form_submit_button(label="送信")

    if submit_button:
        with grpc.insecure_channel(user_service) as channel:
            stub = user_pb2_grpc.UserServiceStub(channel)
            response = stub.Create(
                user_pb2.CreateReqeust(user=user_pb2.NewUser(name=name))
            )
        st.write(response)


run()
