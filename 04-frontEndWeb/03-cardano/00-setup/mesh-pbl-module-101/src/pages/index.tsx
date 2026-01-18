import { useState } from "react";
import Head from "next/head";
import { useWallet, CardanoWallet, MeshBadge } from "@meshsdk/react";

export default function Home() {
  const { connected, wallet } = useWallet();
  const [assets, setAssets] = useState<null | any>(null);
  const [loading, setLoading] = useState<boolean>(false);

  async function getAssets() {
    if (wallet) {
      setLoading(true);
      const _assets = await wallet.getAssets();
      setAssets(_assets);
      setLoading(false);
    }
  }

  return (
    <div className="bg-gray-900 w-full text-white text-center">
      <Head>
        <title>Mesh App on Cardano</title>
        <meta name="description" content="A Cardano dApp powered my Mesh" />
      </Head>
      <main className="flex min-h-screen flex-col items-center justify-center p-24">
        <h1 className="text-6xl font-thin mb-20">
          <a href="https://meshjs.dev/" className="text-sky-600">
            Mesh
          </a>{" "}
          Next.js
        </h1>

        <div className="mb-20">
          <CardanoWallet />
        </div>

        {connected && (
          <div className="mb-20 p-8 bg-gray-800 rounded-xl border border-sky-600 w-full max-w-4xl">
            <h2 className="text-3xl font-bold mb-4">Wallet Assets</h2>
            {assets ? (
              <pre className="text-left overflow-auto max-h-96 bg-gray-900 p-4 rounded mt-4">
                <code className="text-green-400 text-sm">
                  {JSON.stringify(assets, null, 2)}
                </code>
              </pre>
            ) : (
              <button
                type="button"
                onClick={() => getAssets()}
                disabled={loading}
                className={`px-6 py-3 rounded-lg font-bold transition ${
                  loading ? "bg-orange-500 cursor-wait" : "bg-sky-600 hover:bg-sky-500"
                }`}
              >
                {loading ? "Fetching Assets..." : "Get Wallet Assets"}
              </button>
            )}
          </div>
        )}

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 content-center justify-around">
          <a
            href="https://meshjs.dev/apis"
            className="bg-gray-800 rounded-xl border border-white hover:scale-105 transition max-w-96 p-5 m-5"
          >
            <h2 className="text-2xl font-bold mb-2">Documentation</h2>
            <p className="text-gray-400">
              Our documentation provide live demos and code samples; great
              educational tool for learning how Cardano works.
            </p>
          </a>

          <a
            href="https://meshjs.dev/guides"
            className="bg-gray-800 rounded-xl border border-white hover:scale-105 transition max-w-96 p-5 m-5"
          >
            <h2 className="text-2xl font-bold mb-2">Guides</h2>
            <p className="text-gray-400">
              Whether you are launching a new NFT project or ecommerce store,
              these guides will help you get started.
            </p>
          </a>

          <a
            href="https://meshjs.dev/smart-contracts"
            className="bg-gray-800 rounded-xl border border-white hover:scale-105 transition max-w-96 p-5 m-5 md:mx-auto lg:mx-5 md:col-span-2 lg:col-span-1"
          >
            <h2 className="text-2xl font-bold mb-2">Smart Contracts</h2>
            <p className="text-gray-400">
              Open-source smart contracts, complete with documentation, live
              demos, and end-to-end source code.
            </p>
          </a>
        </div>
      </main>
      <footer className="p-8 border-t border-gray-300 flex justify-center">
        <MeshBadge isDark={true} />
      </footer>
    </div>
  );
}
