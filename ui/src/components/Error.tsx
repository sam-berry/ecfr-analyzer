export default function Error({ message }: { message: string }) {
  return (
    <div className="w-full rounded bg-red-50 p-4 text-center font-medium text-red-600">
      {message}
    </div>
  );
}
