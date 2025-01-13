export function useTelegram() {
  //@ts-expect-error ignore
  const tg = window.Telegram.WebApp;
  return { tg, user: tg.initDataUnsafe?.user };
}
