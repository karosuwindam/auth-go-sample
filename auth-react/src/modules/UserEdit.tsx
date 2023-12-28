import axios from "axios";
import { error } from "console";
import { stat } from "fs";
import { useState, useContext, useEffect, useReducer } from "react";

interface User {
  message: string;
  data: {
    id: number;
    name: string;
    auth: number;
  };
}

export const getuser = async (id: number) => {
  const baseURL: string = process.env.REACT_APP_API_URL + "/api/v1/";
  const token = localStorage.getItem("token");
  if (token === null) {
    console.log("token error");
    return null;
  }
  axios.defaults.headers.common["Authorization"] = "Bearer " + token;
  return await axios.get(baseURL + "user/" + id);
};

export const GetUser = (id: number) => {
  const result = getuser(id);
  result.then((res) => {
    const json = res?.data;
    return json;
  });
};
