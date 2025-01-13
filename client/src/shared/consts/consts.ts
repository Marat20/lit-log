//@ts-expect-error ignore
export const TG = window.Telegram.WebApp
export const userId = __IS_DEV__ ? 1929324 : TG.initData.user.id