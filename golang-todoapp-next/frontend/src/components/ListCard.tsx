import Link from "next/link";

const ListCard = ({ todo }: any) => {
  return (
    <div className="card bg-base-100 w-96 shadow-xl">
      <div className="card-body">
        <h2 className="card-title">Title: {todo.Title}</h2>
        <div className="card-actions justify-end">
          <button className="btn btn-primary">
            <Link href={`/details/${todo.ID}`}>Details</Link>
          </button>
        </div>
      </div>
    </div>
  );
};

export default ListCard;
