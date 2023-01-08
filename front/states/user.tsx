import { atom } from 'recoil';
import { recoilPersist } from 'recoil-persist';

const { persistAtom } = recoilPersist();

// undefined: まだログイン確認が完了してない状態とする
// null: ログイン確認をした結果、ログインしていなかった状態とする
export const accessTokenState = atom<string>({
    key: 'accessToken',
    default: '',
    effects_UNSTABLE: [persistAtom]
})