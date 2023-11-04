import { createMachine, assign } from "xstate";

export interface MultiStepFormMachineContext {
  websiteType: WebsiteType;
  websiteDetails: WebsiteDetails;
  errorMessage?: string;
}

interface WebsiteDetails {
  Name: string;
}

type WebsiteType = number;


export type MultiStepFormMachineEvent =
  | {
    type: "BACK";
  }
  | {
    type: "Next";
    info: WebsiteDetails | WebsiteType;
  }
  | {
    type: "CONFIRM";
  };

const multiStepFormMachine =
  createMachine<
    MultiStepFormMachineContext,
    MultiStepFormMachineEvent
  >(
    {
      id: "multiStepForm",
      initial: 'choosingType',
      states: {
        choosingType: {
          on: {
            Next: {
              target: "fillingDetails",
              actions: ["assignWebsiteType"],
            },
          },
        },
        fillingDetails: {
          id: "fillingDetails",
          on: {
            BACK: {
              target: "choosingType",
            },
            Next: {
              target: "confirming",
              actions: ["assignDetailsToContext"],
            },
          },
        },
        confirming: {
          onDone: {
            target: "success",
          },
          initial: "idle",
          states: {
            idle: {
              exit: ["clearErrorMessage"],
              on: {
                CONFIRM: "submitting",
                BACK: {
                  target: "#fillingDetails",
                },
              },
            },
            submitting: {
              invoke: {
                src: "submitPayment",
                onDone: {
                  target: "complete",
                },
                onError: {
                  target: "idle",
                  actions: "assignErrorMessageToContext",
                },
              },
            },
            complete: { type: "final" },
          },
        },
        success: {
          type: "final",
        },
      },
    },
    {
      services: { submitPayment: () => () => { } },
      actions: {
        assignDetailsToContext: assign((context, event) => {
          if (event.type !== "Next") return {};
          return {
            websiteDetails: event.info,
          };
        }),
        assignTypeToContext: assign((context, event) => {
          if (event.type !== "Next") return {};
          return {
            websiteType: event.info,
          };
        }),
        assignErrorMessageToContext: assign((context, event: any) => {
          return {
            errorMessage: event.data?.message || "An unknown error occurred",
          };
        }),
        clearErrorMessage: assign({
          errorMessage: undefined,
        }),
      },
    }
  );

export default multiStepFormMachine;
