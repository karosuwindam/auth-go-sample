import axios from "axios";
import { useState,useContext, useEffect } from "react";

export const PostLogin = (name: string, password: string) => {
    const baseURL:string = process.env.REACT_APP_API_URL+ '/api/v1/';
    // const requestOption = {
    //     method: 'POST',
    //     headers: {
    //         'Content-Type': 'application/x-www-form-urlencoded',
    //     },
    //     body: JSON.stringify({"name": name, "password": password}),
    // };
    // fetch(baseURL + 'login', requestOption)
    axios.post(baseURL + 'login', {"name": name, "password": password})
    .then((res)=> {
        console.log(res);
        localStorage.setItem('token', res.data.token);
    }).catch((error) => {
        console.log('通信失敗');
        console.log(error.status);
    });
    // .then((res)=> res.json()
    // ).then((resJson) =>{
    //     localStorage.setItem('token', resJson.token);
    //     // axios.defaults.headers.common['Authorization'] = resJson.token;

    //   console.log(resJson)
    // }).catch((error) => {
    //     console.log('通信失敗');
    //     console.log(error.status);
    // });
};

export const GetLogin = () => {
    const baseURL:string = process.env.REACT_APP_API_URL+ '/api/v1/';
    const token = localStorage.getItem('token');
    // axios.get(baseURL + 'login',{
    //     headers: {
    //         'Authorization': 'Bearer ' + localStorage.getItem('token')
    //     }
    // })
    // const requestOption = {
    //     // url : baseURL + 'login',
    //     method: 'GET',
    //     headers: {
    //         // 'Content-Type': 'application/x-www-form-urlencoded',
    //         'Authorization': 'Bearer ' + localStorage.getItem('token'),
    //     },
    // };
    // // axios.get(baseURL + 'login')
    // // axios(requestOption)
    // fetch(baseURL + 'login', requestOption)
    // .then((res)=> {
    //   console.log(res);
    // }).catch((error) => {
    //     console.log('通信失敗');
    //     console.log(error.status);
    // });

    /////
    // window.fetch(baseURL + 'login', {
    //     method: "GET",
    //     headers: {
    //       "Content-Type": "application/json",
    //       "Authorization": `Bearer ${localStorage.getItem('token')}`,
    //     },
    //   })
    // .then((res)=> {
    //   console.log(res);
    // }).catch((error) => {
    //     console.log('通信失敗');
    //     console.log(error.status);
    // });
    // axios.defaults.headers.common["Authorization"] = "Bearer " + token;
    // axios.get(baseURL + 'login')
    // .then((response) => {
    //   console.log(response);
    // })
    // .catch((err) => console.log(err));
    const getLogin = async () => {
        try {
            axios.defaults.headers.common["Authorization"] = "Bearer " + token;
            // const response = await axios.get(baseURL );
                const response = await axios.get(baseURL + 'login', );
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
        </form>   
    )
};
