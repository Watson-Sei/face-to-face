import { atom } from 'recoil';
import { CurrentUser } from '../types/user';

// undefined: まだログイン確認が完了してない状態とする
// null: ログイン確認をした結果、ログインしていなかった状態とする
export const currentUserState = atom<undefined | null | CurrentUser>({
    key: 'CurrentUser',
    default: undefined,
})