import Link from 'next/link';
import { useRouter } from 'next/router';
import { useEffect } from 'react';
import { useRecoilState, useRecoilValue } from 'recoil';
import { accessTokenState } from '../states/user';
import logo from '../assets/logo.png';
import google_logo from '../assets/google-logo.png';

export default function Login() {

    const router = useRouter();
    const [accessToken, setAccessToken] = useRecoilState(accessTokenState);

    useEffect(() => {
        if (router.query.code) {
            const code = router.query.code;
            // バックエンドにリクエストを送る
            fetch(`http://localhost:3000/api/auth/token`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    code: code
                })
            })
            .then(res => res.json())
            .then(data => {
                console.log(data)
                setAccessToken(data.access_token)
                router.push('/service')
            })
        }
    }, [router, setAccessToken])

    const handleSignIn = () => {
        const scope = 'https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email';
        const client_id = process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID;
        const redirect_uri = "http://localhost:3001/login"
        const url = `https://accounts.google.com/o/oauth2/v2/auth?scope=${scope}&include_granted_scopes=true&response_type=code&redirect_uri=${redirect_uri}&client_id=${client_id}&access_type=offline`;
        window.location.href = url;
    }

  return (
    <div className='relative h-screen py-16 bg-gradient-to-br from-sky-50 to-gray-200'>
        <div className='relative top-[50%] left-0 translate-y-[-50%] transform container m-auto px-6 text-gray-500 md:px-12 xl:px-40'>
            <div className='m-auto md:w-8/12 lg:w-6/12 xl:w-6/12'>
                <div className='rounded-xl bg-white shadow-xl'>
                    <div className='p-6 sm:p-16'>
                        <div className='space-y-4'>
                            <img src={logo.src} loading="lazy" className='w-10' alt="face-to-face logo" />
                            <h2 className='mb-8 text-2xl text-cyan-900 font-bold'>
                                Sigin in unlock the best of Face To Face
                            </h2>
                        </div>
                        <div className='mt-16 grid space-y-4'>
                            <button onClick={handleSignIn} className='group h-12 px-6 border-2 border-gray-300 rounded-full transition duration-300 hover:border-blue-400 focus:bg-blue-40 active:bg-blue-100'>
                                <div className='relative flex items-center space-x-4 justify-center'>
                                    <img src={google_logo.src} className="absolute left-0 w-5" alt="google logo" />
                                    <span className='block x-max font-semibold tracking-wide text-gray-700 text-sm transition duration-300 group-hover:text-blue-600 sm:text-base'>Continue with Google</span>
                                </div>
                            </button>
                        </div>

                        <div className="mt-32 space-y-4 text-gray-600 text-center sm:-mb-8">
                            <p className="text-xs">By proceeding, you agree to our <a href="#" className="underline">Terms of Use</a> and confirm you have read our <a href="#" className="underline">Privacy and Cookie Statement</a>.</p>
                            <p className="text-xs">This site is protected by reCAPTCHA and the <a href="#" className="underline">Google Privacy Policy</a> and <a href="#" className="underline">Terms of Service</a> apply.</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
  )
}