import axios from "axios";
import { Link } from "react-router-dom";
import "../public/css/Form.css";

import { useState, useContext, useEffect, FormEventHandler } from "react";

interface UserRerister {
  name: string;
  password: string;
  authority: number;
}

const putCreateUser = async (data: UserRerister) => {
  const baseURL: string = process.env.REACT_APP_API_URL + "/api/v1/";
  const token = localStorage.getItem("token");
  if (token === null) {
    console.log("token error");
    return null;
  }
  axios.defaults.headers.common["Authorization"] = "Bearer " + token;
  return await axios.put(baseURL + "user", data);
};

export const UserAdd: React.FC = () => {
  const datalist = {
    auth: [
      { id: 0, name: "guest" },
      { id: 1, name: "user" },
      { id: 2, name: "admin" },
    ],
  };

  const [name, setName] = useState("");
  const [password, setPassword] = useState("");
  const [authority, setAuthority] = useState(0);

  const data = JSON.parse(JSON.stringify(datalist));
  const mAuth = Object.keys(data.auth).map((key) => data.auth[key]);
  const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value);
  };
  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };
  const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setAuthority(Number(e.target.value));
  };
  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const jsondata: UserRerister = {
      name: name,
      password: password,
      authority: authority,
    };
    // console.log(JSON.stringify(jsondata));
    const result = putCreateUser(jsondata);
    result
      .then((res) => {
        setName("");
        setPassword("");
        window.alert("ユーザー登録成功");
        console.log(res);
      })
      .catch((error) => {
        window.alert("ユーザー登録失敗");
        console.log(error);
      });
  };

  return (
    <div className="container">
      <form onSubmit={handleSubmit}>
        <div className="spacer" />
        <label>名前：</label>
        <input
          type="text"
          onChange={handleNameChange}
          name=""
          id=""
          value={name}
        />
        <label>パスワード：</label>
        <input
          type="password"
          onChange={handlePasswordChange}
          name=""
          id=""
          value={password}
        />
        <label>権限：</label>
        <select onChange={(e) => handleChange(e)}>
          {mAuth.map((auth) => (
            <option value={auth.id}>{auth.name}</option>
          ))}
        </select>
        <input type="submit" value="追加" />
      </form>
    </div>
  );
};
