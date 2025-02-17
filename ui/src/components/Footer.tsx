import { IconCopyright } from "@tabler/icons-react";

export default function Footer() {
  return (
    <div className="bg-primary border-primary-700 mt-28 border-t-4 px-4">
      <div className="font-title text-primary-700 pb-10 pt-16 text-center text-3xl uppercase md:pb-16 md:pt-20 md:text-5xl">
        Code of Federal Regulations
      </div>
      <div className="text-primary-700 mt-4 flex items-center justify-center gap-2 pb-4 text-xs">
        <div className="flex items-center gap-1">
          <IconCopyright stroke={1} size={12} /> 2025
        </div>
        <div>|</div>
        <a
          className="hover:text-accent transition ease-in-out"
          target="_blank"
          href="https://github.com/sam-berry/ecfr-analyzer"
        >
          GitHub
        </a>
      </div>
    </div>
  );
}
