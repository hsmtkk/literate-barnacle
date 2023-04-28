import datetime
import os

import booking_pb2
import booking_pb2_grpc
import grpc
import room_pb2
import room_pb2_grpc
import streamlit as st
import user_pb2
import user_pb2_grpc


def run():
    page = st.sidebar.selectbox("Choose your page", ["bookings", "rooms", "users"])
    st.title("APIテスト画面 ユーザー")
    handlers = {
        "bookings": show_bookings_page,
        "rooms": show_rooms_page,
        "users": show_users_page,
    }
    handlers[page]()


def show_bookings_page() -> None:
    booking_service = os.environ["BOOKING"]

    with st.form(key="booking"):
        user_id: int = st.number_input("")
        room_id: int = st.number_input("")
        reserved_num: int = st.number_input("予約人数", step=1)
        date: datetime.date = st.date_input("日付", min_value=datetime.date.today())
        start_time = st.time_input("開始時刻", value=datetime.time(hour=9, minute=0))
        end_time = st.time_input("終了時刻", value=datetime.time(hour=20, minute=0))
        submit_button = st.form_submit_button(label="送信")

        start_date_time = datetime.datetime(
            year=date.year,
            month=date.month,
            day=date.day,
            hour=start_time.hour,
            minute=start_time.minute,
        ).isoformat()
        end_date_time = datetime.datetime(
            year=date.year,
            month=date.month,
            day=date.day,
            hour=end_time.hour,
            minute=end_time.minute,
        ).isoformat()

    if submit_button:
        with grpc.insecure_channel(booking_service) as channel:
            stub = booking_pb2_grpc.BookingServiceStub(channel)
            response = stub.Create(
                booking_pb2.CreateReqeust(
                    booking=booking_pb2.NewBooking(
                        user_id=user_id,
                        room_id=room_id,
                        reserved_num=reserved_num,
                        start_date_time=start_date_time,
                        end_date_time=end_date_time,
                    )
                )
            )
        st.write(response)


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
