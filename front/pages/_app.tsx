import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { useEffect } from 'react';
import { RecoilRoot, useSetRecoilState } from 'recoil';
import { currentUserState } from '../states/user';

export default function App({ Component, pageProps, router }: AppProps) {
  
  return (
    <RecoilRoot>
      <Component {...pageProps} />
    </RecoilRoot>
  )
}
