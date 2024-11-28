import { useEffect, useState } from "react";
import { createSwapy } from "swapy";

const Weight = () => {
  const [data, setData] = useState<Response | null>(null);

  useEffect(() => {
    fetch("http://localhost:8090/test")
      .then((response) => response.json())
      .then((data) => setData(data));
  }, []);

  useEffect(() => {
    const container = document.querySelector(".container")!;
    const swapy = createSwapy(container, {
      swapMode: "hover",
    });
    swapy.onSwap(({ data }) => {
      console.log("swap", data);
      localStorage.setItem("slotItem", JSON.stringify(data.object));
    });

    swapy.onSwapEnd(({ data, hasChanged }) => {
      console.log(hasChanged);
      console.log("end", data);
    });

    swapy.onSwapStart(() => {
      console.log("start");
    });

    return () => {
      swapy.destroy();
    };
  }, []);

  return (
    <div>
      <>
        <div className=" border-[1px] border-neutral-200 mb-2 rounded-md p-2 container">
          <div className="pl-10 w-full flex">
            <div
              className="w-[200px] h-[200px] borer-[1px] border-neutral-200 rounded-md"
              data-swapy-slot="foo"
            >
              xxxx
            </div>
            <div className="w-5/12 pt-[400px]">
              <img src="/public/waga.png" className="w-full" />
            </div>
            <div className="grow ">
              <div className="w-[700px] flex flex-wrap align-middle items-center  ">
                {[
                  "banana.png",
                  "chery.png",
                  "orange.png",
                  "orange2.png",
                  "pinaple.png",
                ].map((fruit) => (
                  <div
                    className="w-[200px] children:h-[200px] "
                    data-swapy-item={fruit}
                    key={fruit}
                  >
                    <img src={`/public/fruits/${fruit}`} />
                  </div>
                ))}
              </div>

              <div className="text-lg">Top 3</div>
              <div className="flex">
                <div className="flex w-1/3">
                  <div className="w-8/12">śliwka</div>
                  <div>3.15 pln</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </>
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
