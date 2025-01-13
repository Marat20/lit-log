import { URL } from "@/shared/api/api";
import { userId } from "@/shared/consts/telegram";
import { Book } from "../model/types/book";

interface Data {
  currentBook?: Book;
  pagesReadToday?: number;
  error?: boolean;
}

export const fetchBook = async (): Promise<Data | undefined> => {
  try {
    const response = await fetch(`${URL}/${userId}/getCurrentBook`);
    const result: Data = await response.json();

    return result;
  } catch (error) {
    console.error(error);
  }
};
