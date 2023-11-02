import axios from "axios";
import { Link } from "react-router-dom";
import { PostLogin,GetLogin, PostLogout } from "./Login"

import { useState,useContext, useEffect } from "react";


export default function GetLoginFunc() {
    return false;
};


export const LoginPage = () => {
    const [form, setForm] = useState({
        name: '',
        password: ''
    });
    const handleForm = (e: any) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value
        });
    }
    return (
        <form>
        <div>
            <label htmlFor="name">名前</label>
            <input id="name" name="name" type="text" onChange={handleForm}/>
            <label htmlFor="password">パスワード</label>
            <input id="password" name="password" type="password" onChange={handleForm}/>
            <button type="button" onClick={()=>PostLogin(form.name, form.password)}>送信</button>
        </div>
        </form>
    )
}

export const LogoutPage = () => {
    return (
        <div>
            <button type="button" onClick={PostLogout}>Logout</button>
        </div>
    
    )
};

export const TestSend = () => {
    const baseURL:string = process.env.REACT_APP_API_URL+ '/api/v1/';
    const token = localStorage.getItem('token');
    const getLogin = async () => {
        try {
            axios.defaults.headers.common["Authorization"] = "Bearer " + token;
                const response = await axios.get(baseURL);
            console.log(response);
        }catch (error){
            console.log(error);
        }
    };
    getLogin();

}

export const Sandbox = () => {
    const [form, setForm] = useState({
        name: '',
        password: ''
    });
    const handleForm = (e: any) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value
        });
    }
    const SendForm = () => {
        console.log(form);
        PostLogin(form.name, form.password);
        
    }
    const GetForm = () => {
        GetLogin();
    }
    
    const TestForm = () => {
        TestSend();
    }
    return (
        <div>
            <form>
                <div>
                    <label htmlFor="name">名前</label>
                    <input id="name" name="name" type="text" onChange={handleForm}/>
                </div>
                <div>
                    <label htmlFor="password">パスワード</label>
                    <input id="password" name="password" type="password" onChange={handleForm}/>
                </div>
                <div>
                    <button type="button" onClick={SendForm}>送信</button>
                </div>
                <div>
                    <button type="button" onClick={GetForm}>GET</button>
                </div>
                <div>
                    <button type="button" onClick={PostLogout}>Logout</button>
                </div>
                <div>
                    <button type="button" onClick={TestForm}>TEST Send</button>
                </div>
            </form>
        <hr />
        <div>
            {/* <button type="button" onClick={Test}>Test</button> */}
            <li><Link to="/">Top</Link></li>
        </div>
        <hr />
        <LoginPage />
        </div>
    )
};
