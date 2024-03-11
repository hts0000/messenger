import axios from "axios";
import camelcaseKey from "camelcase-keys";

import { auth } from "../gen/v1/auth/auth_pb";
import { Messenger } from "./request";
import toast from "react-hot-toast";

export namespace AuthService {
  const url = "http://localhost:18080/v1";

  const instance = axios.create({
    baseURL: url,
    timeout: 3000,
  });

  instance.interceptors.response.use(
    (resp) => {
      return camelcaseKey(resp.data, {
        deep: true,
        pascalCase: true,
      });
    },
    (error) => {
      console.log(error);
      toast.error(error.response.data.message);
      return Promise.reject(error);
    }
  );

  export async function Login(data: auth.v1.IAuthRequest) {
    return Messenger.SendRequest<auth.v1.IAuthRequest, auth.v1.IAuthResponse>({
      method: "POST",
      path: "/auth/login",
      data: data,
    });
  }

  export async function Register(data: auth.v1.IAuthRequest) {
    return Messenger.SendRequest<auth.v1.IAuthRequest, auth.v1.IAuthResponse>({
      method: "POST",
      path: "/auth/register",
      data: data,
    });
  }
}
