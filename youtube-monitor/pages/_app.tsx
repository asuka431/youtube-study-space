import "../styles/global.sass";
import { AppProps } from "next/app";
import createStore from "../store/createStore";
import { Provider } from "react-redux";

export default function App({ Component, pageProps }: AppProps) {
  if (process.env.NEXT_PUBLIC_API_KEY === undefined) {
    console.error('環境変数NEXT_PUBLIC_API_KEYが定義されていません')
  }
  return (
    <Provider store={createStore()}>
      <Component {...pageProps} />
    </Provider>
  );
}
