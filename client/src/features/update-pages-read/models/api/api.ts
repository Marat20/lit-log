import { Book } from "@/entities/book";
import { URL } from "@/shared/api/api";

interface Data {
  pagesRead: number;
  id: string;
}

export interface ReturnData {
  book: Book;
  pagesReadToday: number;
}

export const fetchUpdatePagesRead = async (data: Data) => {
  const { pagesRead, id } = data;
  try {
    const response = await fetch(`${URL}/${id}`, {
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
