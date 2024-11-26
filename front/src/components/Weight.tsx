import { useEffect, useState } from "react";

const Weight = () => {
  const [data, setData] = useState<Response | null>(null);

  useEffect(() => {
    fetch("http://localhost:8090/test")
      .then((response) => response.json())
      .then((data) => setData(data));
  }, []);

  return (
    <div>
      {data != null && (
        <>
          {data.map((e) => {
            return (
              <div className=" border-[1px] border-neutral-200 mb-2 rounded-md p-2">
                <div className="mb-2">{e.file.split("/").pop()}</div>
                <div className="flex mb-5">
                  <div className="w-5/12">
                    <img
                      src={"http://localhost:8090/file?file=" + e.file}
                      className="max-w-full"
                    />
                  </div>
                  <div className="pl-10 w-full">
                    {e.probs.map((p) => {
                      return (
                        <div className="flex">
                          <div className="w-8/12">{p.label}</div>
                          <div>{p.confidence}</div>
                        </div>
                      );
                    })}
                  </div>
                  <div>
                    <img
                      src={
                        "http://localhost:8090/example?label=" +
                        e.probs[0].label
                      }
                      className="max-w-full"
                    />
                  </div>
                </div>
              </div>
            );
          })}
        </>
      )}
    </div>
  );
};

export { Weight };

type Response = IObject[];

interface IObject {
  file: string;
  probs: Prob[];
}

interface Prob {
  label: string;
  confidence: number;
}
