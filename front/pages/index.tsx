import useRequireLogin from "../hooks/useRequireLogin"
import bg_top from '../assets/lp/bg-top.svg';

export default function LP() {

  return (
    <div className=" pr-[1.5rem] pl-[1.5rem]">
      <div className="pt-[5rem] pb-[5rem]">
        {/* catch phrase */}
        <div className="text-center leading-[1.4]">
          <h1 className="md:text-[5.5rem] text-[3.5rem] font-semibold mb-5">Make your interview 
          <span className="block top-font">Wonderful</span></h1>
          <div>
            <p className="text-[1.25rem] opacity-1 text-[#666666] mb-5">Third-party recording and openness of interview data improves the quality of media sources.</p>
          </div>
        </div>
        {/* two button */}
        <div className="flex justify-center mb-5">
          <button type="button" className="mr-2 px-5 py-3 text-base font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
            Start free trial
          </button>
          <button type="button" className="ml-2 px-5 py-3 text-base font-medium text-center text-white bg-gray-800 rounded-lg hover:bg-gray-900 focus:ring-4 focus:outline-none focus:ring- dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
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
    </div>
  )
}