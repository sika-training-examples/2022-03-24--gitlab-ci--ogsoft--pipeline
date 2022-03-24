import React from "react";
import Head from "next/head";

export default function Home() {
  return (
    <>
      <Head>
        <title>Hello OG Soft!</title>
      </Head>
      <h1>
        Hello <span style={{ color: "orange" }}>O</span>
        <span style={{ color: "green" }}>G</span> Soft!
      </h1>
    </>
  );
}
