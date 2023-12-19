import axios from "axios";
import { ViewPage } from "./ViewPage";
import { useEffect, useState } from "react";

export const UserList = () => {
  const [data, setData] = useState([]);
  //1秒ごとにユーザー一覧情報を取得する
  const [count, setCount] = useState(0);
  const [delay, setDelay] = useState(1000);

  const getUserList = () => {
    const baseURL: string = process.env.REACT_APP_API_URL + "/api/v1/";
    const token = localStorage.getItem("token");
    if (token === null) {
      console.log("token error");
    }
    axios.defaults.headers.common["Authorization"] = "Bearer " + token;
    axios
      .get(baseURL + "user/list")
      .then((res) => {
        const json = res.data;
        for (let i = 0; i < json.data.length; i++) {
          //authの値からroleを取得
          if (json.data[i].auth == 0) {
            json.data[i].role = "guest";
          } else if (json.data[i].auth == 1) {
            json.data[i].role = "user";
          } else if (json.data[i].auth >= 2) {
            json.data[i].role = "admin";
          }
        }
        //UserListの宣言
        setData(json.data);
      })
      .catch((error) => {
        console.log("通信失敗");
        console.log(error.status);
        localStorage.removeItem("token");
      });
  };

  useEffect(() => {
    const interval = setInterval(() => {
      getUserList();
      setCount(count + 1);
    }, delay);
    return () => clearInterval(interval);
  }, []);
  return (
    <ViewPage roles={["admin"]}>
      <h1>User List</h1>
      {/*dataがnull以外の時 */}
      {data && (
        <table>
          <thead>
            <tr>
              <th>id</th>
              <th>name</th>
              <th>role</th>
              <th></th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {data.map((user: any) => (
              <tr key={user.id}>
                <td>{user.id}</td>
                <td>{user.name}</td>
                <td>{user.role}</td>
                <th>view</th>
                <th>edit</th>
              </tr>
            ))}
          </tbody>
        </table>
      )}
      {/*dataがnullの時 */}
      {!data && <p>loading...</p>}
    </ViewPage>
  );
};

export const Edit = () => {
  const [data, setData] = useState([]);
  //1秒ごとにユーザー一覧情報を取得する
  const [count, setCount] = useState(0);
  const [delay, setDelay] = useState(1000);

  const getUserList = () => {
    const baseURL: string = process.env.REACT_APP_API_URL + "/api/v1/";
    const token = localStorage.getItem("token");
    if (token === null) {
      console.log("token error");
    }
    axios.defaults.headers.common["Authorization"] = "Bearer " + token;
    axios
      .get(baseURL + "user/list")
      .then((res) => {
        const json = res.data;
        for (let i = 0; i < json.data.length; i++) {
          //authの値からroleを取得
          if (json.data[i].auth == 0) {
            json.data[i].role = "guest";
          } else if (json.data[i].auth == 1) {
            json.data[i].role = "user";
          } else if (json.data[i].auth >= 2) {
            json.data[i].role = "admin";
          }
        }
        //UserListの宣言
        setData(json.data);
      })
      .catch((error) => {
        console.log("通信失敗");
        console.log(error.status);
        localStorage.removeItem("token");
      });
  };

  useEffect(() => {
    const interval = setInterval(() => {
      getUserList();
      setCount(count + 1);
    }, delay);
    return () => clearInterval(interval);
  }, []);
  return (
    <ViewPage roles={["admin"]}>
      <h1>Edit</h1>
      {/*dataがnull以外の時 */}
      {data && (
        <table>
          <thead>
            <tr>
              <th>id</th>
              <th>name</th>
              <th>role</th>
            </tr>
          </thead>
          <tbody>
            {data.map((user: any) => (
              <tr key={user.id}>
                <td>{user.id}</td>
                <td>{user.name}</td>
                <td>{user.role}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
      {/*dataがnullの時 */}
      {!data && <p>loading...</p>}
    </ViewPage>
  );
};
