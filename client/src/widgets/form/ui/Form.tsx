import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';

const schema = z.object({
  title: z.string(),
  author: z.string(),
  totalPages: z.number(),
  dailyGoal: z.number(),
}).required();

type Schema = z.infer<typeof schema>;

export const Form = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Schema>({ resolver: zodResolver(schema) });

  const onSubmit = handleSubmit((data) => console.log(data));

  return (
    <form onSubmit={onSubmit}>
      <input {...register('title')} placeholder='Title' />
      {errors?.title && <p>{errors.title.message}</p>}
      <input {...register('author')} placeholder='Author' />
      {errors?.author && <p>{errors.author.message}</p>}
      <input {...register('totalPages')} placeholder='TotalPages' />
      {errors?.totalPages && <p>{errors.totalPages.message}</p>}
      <input {...register('dailyGoal')} placeholder='DailyGoal' />
      {errors?.dailyGoal && <p>{errors.dailyGoal.message}</p>}
      <input type='submit' />
    </form>
  );
};
