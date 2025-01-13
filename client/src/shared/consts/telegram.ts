//@ts-expect-error ignore
const tg = window.Telegram.WebApp;
export const userId = __IS_DEV__ ? 1929324 : tg.initDataUnsafe?.user.id;
