"use client";

import { useRouter } from "next/navigation";
import { FormEvent } from "react";

const getDeleteTodo = async (id: string) => {
  const res = await fetch(`http://localhost:8000/todos/${id}`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    throw new Error("Failed to get a single post todo");
  }
  return res.json();
};

const DeleteTodo = ({ todo }: any) => {
  const router = useRouter();

  const handleDelete = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    await getDeleteTodo(todo.ID);

    router.push("/");
  };

  return (
    <div>
      <form
        onSubmit={async (e) => {
          await handleDelete(e);
        }}
      >
        <button className="btn btn-primary w-full gap-5">Delete</button>
      </form>
    </div>
  );
};

export default DeleteTodo;
