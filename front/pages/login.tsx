import Link from 'next/link';
import { useRouter } from 'next/router';
import { useEffect } from 'react';

export default function Login() {

    const router = useRouter();

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
            })
        }
    }, [router])

    const handleSignIn = () => {
        const scope = 'https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email';
        const client_id = process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID;
        const redirect_uri = "http://localhost:3001/login"
        const url = `https://accounts.google.com/o/oauth2/v2/auth?scope=${scope}&include_granted_scopes=true&response_type=code&redirect_uri=${redirect_uri}&client_id=${client_id}&access_type=offline`;
        window.location.href = url;
    }

  return (
    <div>
        <button onClick={handleSignIn}>Sign in with Google</button>
    </div>
  )
}