export default function Navigation({
  title = "Code of Federal Regulations",
}: {
  title?: string;
}) {
  return (
    <div className="font-title bg-primary text-light px-4 py-2 text-center text-3xl uppercase tracking-wide">
      {title}
    </div>
  );
}
