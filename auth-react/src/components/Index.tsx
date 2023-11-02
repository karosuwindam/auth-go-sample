
import logo from './img/logo.svg';
import './css/Index.css'
import { Link } from "react-router-dom";
// import { PostLogin, GetLogin, PostLogout } from './Login';
import React, {useState, useReducer , useEffect } from "react";
import { LoginPage,GetLogin } from './Login';

export const Index = () =>{
    // //localStorage内にあるtoken情報が更新した場合に、再レンダリングする
    // const [token, setToken] = useReducer((state: any, newState: any) => ({...state, ...newState}), localStorage.getItem('token'));
    // useEffect(() => {
    //     setToken(localStorage.getItem('token'));
    // }, [localStorage.getItem('token')]);


    return (
        <div className="App">
          <div className="header">
            {/*変数情報が含んだLoginPageを表示させる*/}
            {/* localstorageのtokenがnull出ない場合は、LoginPage */}
            {/* nullの場合は、LogoutPage */}
            {/* {token === null ? <LoginPage /> : <LogoutPage />} */}
            <LoginPage />
            
          </div>
          <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <Link to="/sandbox" className='App-link'>Sandbox Page</Link>
          </header>
        </div>
    )
};

