import Toolbar from '../components/toolbar';

function UserPage() {
  return (
    <div className="h-full">
      <Toolbar />
      <ul className="absolute h-[calc(100%-45px)] w-0 overflow-hidden bg-gray-200 duration-300" id="lateral-menu">
        <li className="flex h-[35px] w-full cursor-pointer select-none items-center whitespace-nowrap border-b border-zinc-500 bg-zinc-300 pl-4 font-nunito tracking-wider hover:bg-zinc-400">Products</li>
        <li className="flex h-[35px] w-full cursor-pointer select-none items-center whitespace-nowrap border-b border-zinc-500 bg-zinc-300 pl-4 font-nunito tracking-wider hover:bg-zinc-400">Services</li>
        <li className="flex h-[35px] w-full cursor-pointer select-none items-center whitespace-nowrap border-b border-zinc-500 bg-zinc-300 pl-4 font-nunito tracking-wider hover:bg-zinc-400">Our policies</li>
      </ul>
    </div>
  );
}

export default UserPage;
