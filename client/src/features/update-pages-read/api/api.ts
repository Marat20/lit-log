import { Book } from "@/entities/book";
import { URL } from "@/shared/api/api";
import { userId } from "@/shared/consts/telegram";

interface Data {
  pagesRead: number;
  bookId: string;
}

export interface ReturnData {
  currentBook: Book;
  pagesReadToday: number;
}

export const fetchUpdatePagesRead = async (data: Data) => {
  const { pagesRead, bookId } = data;
  try {
    const response = await fetch(`${URL}/${userId}/${bookId}`, {
      method: "PUT",
      body: JSON.stringify({ pagesRead }),
    });
    const result: ReturnData = await response.json();

    return result;
  } catch (error) {
    console.error(error);
    throw error;
  }
};
