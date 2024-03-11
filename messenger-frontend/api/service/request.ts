import axios from "axios";
import toast from "react-hot-toast";
import camelcaseKey from "camelcase-keys";

export namespace Messenger {
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

  const WithAuth = instance.interceptors.request.use(
    (config) => {
      // add token
      return config;
    },
    (error) => {
      return Promise.reject(error);
    }
  );

  export interface RequestOption<REQ> {
    method: "GET" | "POST" | "PUT" | "DELETE";
    path: string;
    data?: REQ;
  }

  export async function SendRequest<REQ, RES>(o: RequestOption<REQ>) {
    // 屏蔽添加token
    instance.interceptors.request.eject(WithAuth);
    return instance<RES>({
      method: o.method,
      url: o.path,
      data: o.data,
    });
  }

  export async function SendRequestWithAuth<REQ, RES>(o: RequestOption<REQ>) {
    return instance<RES>({
      method: o.method,
      url: o.path,
      data: o.data,
    });
  }
}
