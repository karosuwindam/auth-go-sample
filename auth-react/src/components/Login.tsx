import axios from "axios";
import { useState,useContext, useEffect } from "react";

export const PostLogin = (name: string, password: string) => {
    const baseURL:string = process.env.REACT_APP_API_URL+ '/api/v1/';
    
    axios.post(baseURL + 'login', {"name": name, "password": password})
    .then((res)=> {
        console.log(res);
        const json = res.data;
        localStorage.setItem('token', json.data.token);
        axios.defaults.headers.common["Authorization"] = "Bearer " + json.data.token;
    }).catch((error) => {
        console.log('通信失敗');
        console.log(error.status);
    });
};

export const GetLogin = () => {
    const baseURL:string = process.env.REACT_APP_API_URL+ '/api/v1/';
    const token = localStorage.getItem('token');
    if (token === null){
        console.log('token error');
        return false;
    }
    axios.defaults.headers.common["Authorization"] = "Bearer " + token;
    axios.get(baseURL + 'login', )
    .then((res)=>{
        const json = res.data;
        console.log(res);
        localStorage.setItem('token', json.data.token);
    }).catch((error) =>{
        console.log('通信失敗');
        console.log(error.status);
        localStorage.removeItem('token');
    });
}

export default function GetLoginFunc() {
    return false;
};

export const PostLogout = () => {
    const baseURL:string = process.env.REACT_APP_API_URL+ '/api/v1/';
       
    axios.post(baseURL + 'logout')
    .then((res)=> {
        console.log(res);
    }).catch((error) => {
        console.log('通信失敗');
        console.log(error.status);
    });
    localStorage.removeItem('token');
    delete axios.defaults.headers.common["Authorization"];

}

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

export const Login = () => {
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
            <hr />
            <div>
                <button type="button" onClick={TestForm}>TEST Send</button>
            </div>
        </form>
    )
};
