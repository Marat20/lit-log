import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import clsx from "clsx";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import { z } from "zod";
import { addNewBook } from "../api/api";
import cls from "./form.module.scss";

const schema = z
  .object({
    title: z
      .string({
        required_error: "Title is required",
        invalid_type_error: "Title must be a string",
      })
      .min(1, { message: "Must be 1 or more characters long" }),
    author: z
      .string({
        required_error: "Author is required",
        invalid_type_error: "Author must be a string",
      })
      .min(1, { message: "Must be 1 or more characters long" }),
    totalPages: z
      .number({
        required_error: "Total pages is required",
        invalid_type_error: "Total pages must be a number",
      })
      .gte(1, { message: "Must be 1 or more characters long" }),
    dailyGoal: z
      .number({
        required_error: "Daily goal is required",
        invalid_type_error: "Daily goal must be a number",
      })
      .gte(1, { message: "Must be 1 or more characters long" }),
  })
  .required();

type Schema = z.infer<typeof schema>;

export const Form = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Schema>({ resolver: zodResolver(schema) });

  const navigate = useNavigate();

  const queryClient = useQueryClient();

  const { mutate } = useMutation({
    mutationFn: addNewBook,
    onSuccess: (data) => {
      queryClient.setQueryData(["book"], {
        book: data.book,
        pagesReadToday: data.pagesReadToday,
      });
      navigate(`/${data.book.id}`);
    },
  });

  const onSubmit = handleSubmit((data) => {
    mutate({ ...data });
  });

  return (
    <form onSubmit={onSubmit} className={cls.form}>
      <section className={clsx(errors?.title && cls.section)}>
        <input {...register("title")} placeholder='Title' type='text' />
      </section>
      {/* {errors?.title && <ErrorMessage>{errors.title.message}</ErrorMessage>} */}

      <section className={clsx(errors?.author && cls.section)}>
        <input {...register("author")} placeholder='Author' type='text' />
      </section>
      {/* {errors?.author && <ErrorMessage>{errors.author.message}</ErrorMessage>} */}

      <section className={clsx(errors?.totalPages && cls.section)}>
        <input
          {...register("totalPages", { valueAsNumber: true })}
          placeholder='Total Pages'
          type='number'
        />
      </section>
      {/* {errors?.totalPages && (
        <ErrorMessage>{errors.totalPages.message}</ErrorMessage>
      )} */}

      <section className={clsx(errors?.dailyGoal && cls.section)}>
        <input
          {...register("dailyGoal", { valueAsNumber: true })}
          placeholder='Daily Goal'
          type='number'
        />
      </section>
      {/* {errors?.dailyGoal && (
        <ErrorMessage>{errors.dailyGoal.message}</ErrorMessage>
      )} */}

      <input className={cls.btn_submit} type='submit' />
    </form>
  );
};
