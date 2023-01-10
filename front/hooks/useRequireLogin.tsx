import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import { useRecoilState } from "recoil";
import { accessTokenState } from "../states/user";

export default function useRequireLogin(role: string) {
    const router = useRouter();

    const [loading, setLoading] = useState(true)
    const [accessToken, setAccessToken] = useRecoilState(accessTokenState);

    useEffect(() => {
        if (accessToken) {
            // トークンの有効性を確認する
            const headers = {
                'Authorization': `Bearer ${accessToken}`
            }
            
            fetch(`http://localhost:3000/api/${role}/check`, {headers})
            .then(res => {
                if (res.ok) {
                    setLoading(false)
                    return;
                } else {
                    setAccessToken("")
                    router.push('/login')
                }
            })
            .catch(_=> {
                setAccessToken("")
                router.push('/login')
            })
        } else {
            router.push('/login')
        }
    }, [accessToken, role, router, setAccessToken])

    return {
        loading,
    }
}