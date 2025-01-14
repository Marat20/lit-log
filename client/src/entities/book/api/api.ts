import { URL } from "@/shared/api/api";
import { userId } from "@/shared/consts/telegram";
import { Book } from "../model/types/book";

interface FetchBookReturnData {
  book?: Book;
  pagesReadToday?: number;
  error?: string;
}

export const fetchBook = async (
  bookId?: string,
): Promise<FetchBookReturnData | undefined> => {
  try {
    const response = await fetch(`${URL}/${userId}/${bookId}`);
    const result = await response.json();

    return result;
  } catch (error) {
    console.error(error);
  }
};

interface InitReturnData {
  bookId: string;
}

export const init = async (): Promise<InitReturnData | undefined> => {
  try {
    const response = await fetch(`${URL}/${userId}/init`);
    const result = await response.json();

    return result;
  } catch (error) {
    console.error(error);
  }
};
