import Image from "next/image";
import {Button} from "@/components/ui/button";
import axios from "axios";
import {hello} from "@/api/gen/v1/hello/hello_pb";

export default function Home() {
  axios.get<hello.v1.HelloResponse>("http://localhost:18080/v1/hello/hts0000").then(resp => {
      console.log(resp.data)
  }).catch(err => console.log(err))
  return (
    <main className="">
      Hello world
      <Button>Click me</Button>
    </main>
  );
}
