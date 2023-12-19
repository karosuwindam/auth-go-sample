import axios from "axios";
import { ViewPage } from "./ViewPage";
import { useEffect, useState } from "react";

export const UserType = {
  name: "",
  role: "",
};

export type User = (typeof UserType)[keyof typeof UserType];

export const UserJsonType = {
  name: "",
  auth: 0,
};

export type UserJson = (typeof UserJsonType)[keyof typeof UserJsonType];

export const inputUser = (data: UserJson) => {
  var role = "guest";
  const tmp = {
    name: data,
    role: role,
  };
  return tmp;
};

export const GetUserList = () => {
  const baseURL: string = process.env.REACT_APP_API_URL + "/api/v1/";
  const token = localStorage.getItem("token");
  if (token === null) {
    console.log("token error");
    return null;
  }
  axios.defaults.headers.common["Authorization"] = "Bearer " + token;
  axios
    .get(baseURL + "user/list")
    .then((res) => {
      const json = res.data;
      //UserListの宣言
      var users: User[] = [];
      for (let i = 0; i < json.data.length; i++) {
        users.push(json.data[i]);
      }
      console.log(users);
      return users;
    })
    .catch((error) => {
      console.log("通信失敗");
      console.log(error.status);
      localStorage.removeItem("token");
      return null;
    });
};

export const Edit = () => {
  //1秒ごとにユーザー一覧情報を取得する
  const [count, setCount] = useState(0);
  const [delay, setDelay] = useState(1000);
  useEffect(() => {
    const interval = setInterval(() => {
      GetUserList();
      setCount(count + 1);
    }, delay);
    return () => clearInterval(interval);
  }, []);
  return (
    <ViewPage roles={["admin"]}>
      <h1>Edit</h1>
    </ViewPage>
  );
};
