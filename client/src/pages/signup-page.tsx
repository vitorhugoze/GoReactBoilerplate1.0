import { useEffect, useState } from 'react';
import Toolbar from '../components/toolbar';
import signupHandler from '../handlers/signup-handler';

type InputParams = {
  desc: string;
  id: string;
  type: string;
};

const doSignup = async function (fieldData: Map<string, string>): Promise<boolean> {
  const keys = ['pess-name', 'user-name', 'user-pass', 'user-mail'];
  var user, pess, mail, pass: string;

  for (var k in keys) {
    if (fieldData.get(k) === null) {
      return false;
    }
  }

  pess = fieldData.get(keys[0])!;
  user = fieldData.get(keys[1])!;
  pass = fieldData.get(keys[2])!;
  mail = fieldData.get(keys[3])!;

  var res = await signupHandler({
    user: user,
    pess: pess,
    mail: mail,
    pass: pass,
  });

  return res;
};

function SignupPage() {
  const [fieldValue, setFieldValue] = useState(new Map<string, string>());
  const [created, setCreated] = useState(false);

  useEffect(() => {}, [created]);

  const UserInputs = (params: InputParams) => {
    return (
      <div className="mt-5 flex">
        <h3 className="w-24 font-nunito tracking-wide">{params.desc}</h3>
        <input
          onChange={(e) => {
            setFieldValue(fieldValue.set(params.id, e.target.value));
          }}
          className="h-[30px] w-56 rounded border bg-zinc-300 pl-1 focus:outline-none focus:ring-0"
          type={params.type}
          name=""
          id={params.id}
        />
      </div>
    );
  };

  return (
    <div className="h-full">
      <Toolbar />
      <div className="flex h-full w-full items-center justify-center">
        <div className="flex flex-col justify-center">
          <UserInputs desc="Full name:" id="pess-name" type="text" />
          <UserInputs desc="Username:" id="user-name" type="text" />
          <UserInputs desc="Password:" id="user-pass" type="password" />
          <UserInputs desc="E-mail:" id="user-mail" type="text" />
          <div className="flex">
            <button
              onClick={() => {
                doSignup(fieldValue).then((res) => setCreated(res));
              }}
              className="mt-3 h-[30px] w-20 rounded bg-slate-400"
            >
              Submit
            </button>
            <span className="ml-3 flex items-end font-nunito font-semibold">{created ? 'User created!' : ''}</span>
          </div>
        </div>
      </div>
    </div>
  );
}

export default SignupPage;
