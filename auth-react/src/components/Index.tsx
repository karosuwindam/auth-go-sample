import logo from "./img/logo.svg";
import "./css/Index.css";
import { Link } from "react-router-dom";
// import { PostLogin, GetLogin, PostLogout } from './Login';
import React, { useState, useReducer, useEffect } from "react";
import { LoginPage, GetLogin } from "./Login";

export const Index = () => {
  // //localStorage内にあるtoken情報が更新した場合に、再レンダリングする
  // const [token, setToken] = useReducer((state: any, newState: any) => ({...state, ...newState}), localStorage.getItem('token'));
  // useEffect(() => {
  //     setToken(localStorage.getItem('token'));
  // }, [localStorage.getItem('token')]);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <Link to="/sandbox" className="App-link">
          Sandbox Page
        </Link>
        <Link to="/home" className="App-link">
          Home Page
        </Link>
        <Link to="/page1" className="App-link">
          Page1
        </Link>
      </header>
    </div>
  );
};
