import axios from "axios";

export namespace Messenger {
  const url = "http://localhost:18080/v1";

  const authData = {
    token: "",
    expiryMs: 0,
  };

  export interface RequestOption<REQ> {
    method: "GET" | "POST" | "PUT" | "DELETE";
    path: string;
    data?: REQ;
  }

  export async function SendRequest<REQ, RES>(
    o: RequestOption<REQ>
  ): Promise<RES> {
    return new Promise((resolve, reject) => {
      const instance = axios.create({
        baseURL: url,
        timeout: 3000,
      });

      instance
        .request<RES>({
          method: o.method,
          url: o.path,
          data: o.data,
        })
        .then((resp) => {
          resolve(resp.data);
        })
        .catch((error) => {
          reject(error.response.data.message as string);
        });
    });
  }
}
