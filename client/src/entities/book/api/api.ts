import { URL } from "@/shared/api/api";
import { Book } from "../model/types/book";

export interface FetchBookReturnData {
  book?: Book;
  pagesReadToday?: number;
  error?: string;
}

export const fetchBook = async (bookId?: string) => {
  try {
    const response = await fetch(`${URL}/${bookId}`);
    const result: FetchBookReturnData = await response.json();

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
    const response = await fetch(`${URL}/init`);
    const result = await response.json();

    return result;
  } catch (error) {
    console.error(error);
  }
};
