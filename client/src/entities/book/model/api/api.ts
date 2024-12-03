import { Book } from "../types/book";

export const fetchBook = async (): Promise<Book | undefined> => {
  try {
    const response = await fetch("http://localhost:8080/books/EFMpb0ckgL2hC1giqz9d-");
    const result = await response.json();
    return result;
  } catch (error) {
    console.error(error);
  }
};
