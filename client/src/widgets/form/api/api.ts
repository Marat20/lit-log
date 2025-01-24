import { Book } from "@/entities/book";
import { URL } from "@/shared/api/api";

interface AddNewBookData {
  title: string;
  author: string;
  totalPages: number;
  dailyGoal: number;
}

export interface AddNewBookReturnData {
  book: Book;
  pagesReadToday: number;
}

export const addNewBook = async (data: AddNewBookData) => {
  try {
    const response = await fetch(`${URL}/new`, {
      method: "POST",
      body: JSON.stringify({ ...data }),
    });
    const result: AddNewBookReturnData = await response.json();

    return result;
  } catch (error) {
    console.error(error);
    throw error;
  }
};
