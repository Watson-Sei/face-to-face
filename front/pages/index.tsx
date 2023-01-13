/* eslint-disable @next/next/no-img-element */
import useRequireLogin from "../hooks/useRequireLogin"
import bg_top from '../assets/lp/bg-top.svg';
import bg_pilot from '../assets/lp/bg-pilot.svg';
import { FingerPrintIcon } from '@heroicons/react/24/outline';
import { useRouter } from "next/router";

export default function LP(): JSX.Element {

  const router = useRouter();

  return (
    <div className=" pr-[1.5rem] pl-[1.5rem]">
      <div className="pt-[5rem] pb-[5rem]">
        {/* catch phrase */}
        <div className="text-center leading-[1.4]">
          <h1 className="md:text-[5.5rem] text-[3.5rem] font-semibold mb-5">あなたのインタビューを
          <span className="block top-font">素晴らしく</span></h1>
          <div>
            <p className="mr-auto ml-auto mt-0 text-[1.25rem] opacity-1 text-[#666666] mb-5 max-w-[768px]">
              インタビューの第三者による録画・公開は、メディアのソースの信憑性や捏造といった品質問題を解決します。また、Blockchainを利用し永久的にデータを保存します。
            </p>
          </div>
        </div>
        {/* two button */}
        <div className="flex justify-center mb-5">
          <button onClick={() => router.push("/service")} type="button" className="mr-2 px-5 py-3 text-base font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
            Start free trial
          </button>
          <button onClick={() => router.push("/")} type="button" className="ml-2 px-5 py-3 text-base font-medium text-center text-white bg-gray-800 rounded-lg hover:bg-gray-900 focus:ring-4 focus:outline-none focus:ring- dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
            Learn more
          </button>
        </div>
        {/* top image */}
        <div className="flex justify-center">
          <div className="shadow-lg">
            <img src={bg_top.src} alt="top image" className=" w-[768px] h-[432px]" />
          </div>
        </div>
      </div>
      <div className="pt-[5rem] pb-[5rem]">
        <div className="text-center">
          <h1 className="text-[2.625rem] font-semibold mb-5">
            シンプルな仕組み
          </h1>
          <p className="mr-auto ml-auto mt-0 text-[1.25rem] opacity-1 text-[#666666] mb-5 max-w-[768px]">
          Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur excepteur sint occaecat cupidatat.
          </p>
        </div>
        <div className="flex justify-center">
          <div className="bg-black w-[768px] h-[432px]"></div>
        </div>
      </div>
      <div className="pt-[5rem] pb-[5rem]">
        <div className="text-center">
          <h1 className="text-[2.625rem] font-semibold mb-5">
            ソリューション
          </h1>
          <p className="mr-auto ml-auto mt-0 text-[1.25rem] opacity-1 text-[#666666] mb-5 max-w-[768px]">
          Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur excepteur sint occaecat cupidatat.
          </p>
        </div>
        <div className="pt-5 md:h-[1050px] h-[1600px]">
          <div className="pilot-base w-[300px] h-[300px] md:w-[400px] md:h-[400px]">
            <img src={bg_pilot.src} alt="top image" className="w-[400px] h-auto m-auto absolute" />
            <div className="max-w-[700px] relative md:top-[120%] top-[250%] flex justify-center flex-wrap">
              <div className="md:w-[700px] w-[320px] flex justify-around flex-wrap">
                <div className="w-[320px] h-[200px] shadow-lg bg-white mb-4 p-6">
                  <div className="w-[56px] h-[56px] rounded-[50%] bg-blue-600 m-auto flex justify-center items-center">
                    <FingerPrintIcon className="w-[24px] h-[24px]" color="white" />
                  </div>
                  <div className="text-center">
                    <h1 className="font-semibold pt-3">Headless CMS</h1>
                    <p className="pt-2">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
                  </div>
                </div>
                <div className="w-[320px] h-[200px] shadow-lg bg-white mb-4 p-6">
                  <div className="w-[56px] h-[56px] rounded-[50%] bg-blue-600 m-auto flex justify-center items-center">
                    <FingerPrintIcon className="w-[24px] h-[24px]" color="white" />
                  </div>
                  <div className="text-center">
                    <h1 className="font-semibold pt-3">Headless CMS</h1>
                    <p className="pt-2">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
                  </div>
                </div>
              </div>
              <div className="md:w-[670px] w-[320px] flex justify-around flex-wrap">
              <div className="w-[320px] h-[200px] shadow-lg bg-white mb-4 p-6">
                  <div className="w-[56px] h-[56px] rounded-[50%] bg-blue-600 m-auto flex justify-center items-center">
                    <FingerPrintIcon className="w-[24px] h-[24px]" color="white" />
                  </div>
                  <div className="text-center">
                    <h1 className="font-semibold pt-3">Headless CMS</h1>
                    <p className="pt-2">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
                  </div>
                </div>
                <div className="w-[320px] h-[200px] shadow-lg bg-white mb-4 p-6">
                  <div className="w-[56px] h-[56px] rounded-[50%] bg-blue-600 m-auto flex justify-center items-center">
                    <FingerPrintIcon className="w-[24px] h-[24px]" color="white" />
                  </div>
                  <div className="text-center">
                    <h1 className="font-semibold pt-3">Headless CMS</h1>
                    <p className="pt-2">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
                  </div>
                </div>
              </div>
              <div className="md:w-[670px] w-[320px] flex justify-around flex-wrap">
                <div className="w-[320px] h-[200px] shadow-lg bg-white mb-4 p-6">
                  <div className="w-[56px] h-[56px] rounded-[50%] bg-blue-600 m-auto flex justify-center items-center">
                    <FingerPrintIcon className="w-[24px] h-[24px]" color="white" />
                  </div>
                  <div className="text-center">
                    <h1 className="font-semibold pt-3">Headless CMS</h1>
                    <p className="pt-2">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
                  </div>
                </div>
                <div className="w-[320px] h-[200px] shadow-lg bg-white mb-4 p-6">
                  <div className="w-[56px] h-[56px] rounded-[50%] bg-blue-600 m-auto flex justify-center items-center">
                    <FingerPrintIcon className="w-[24px] h-[24px]" color="white" />
                  </div>
                  <div className="text-center">
                    <h1 className="font-semibold pt-3">Headless CMS</h1>
                    <p className="pt-2">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="pt-[5rem] pb-[5rem]">
        <div className="w-full h-[275px] bg-blue-600 text-center p-10 flex justify-center items-center flex-wrap">
          <div className="w-[100vw]">
            <h3 className="text-[2rem] text-white font-semibold mb-[8px]">
              Ready to get started?
            </h3>
            <p className="text-white opacity-75">We have a generous free tier available to get you started right away.</p>
          </div>
          <div className="w-[100vw]">
            <button onClick={() => router.push("/service")} type="button" className="px-5 py-3 text-base font-medium text-center text-blue-600 bg-white rounded-lg hover:bg-gray focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Get started for free</button>
          </div>
        </div>
      </div>
    </div>
  )
}