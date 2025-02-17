"use client";

import CountUp from "react-countup";

export default function NumberCounter({
  start,
  end,
}: {
  start: number;
  end: number;
}) {
  return <CountUp start={start} end={end}></CountUp>;
}
