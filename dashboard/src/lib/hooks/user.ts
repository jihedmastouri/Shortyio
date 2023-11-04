import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { z } from "zod";

const UserSchema = z.object({
  email: z.string().email(),
  name: z.string(),
  role: z.string(),
});

type User = z.infer<typeof UserSchema>;

const useUser = () => {
  const queryClient = useQueryClient();

  const getUser = async () => {
    return useQuery(["user"], async () => await axios.get("/api/user"));
  };

  const setUser = async (user: User) => {
    return useMutation(["user"], {
      mutationFn: async () => await axios.post("/api/user", user),
      onMutate: async () => {
        await queryClient.cancelQueries(["user"]);
        const previousUser = queryClient.getQueryData(["user"]);
        queryClient.setQueryData(["user"], user);
        return { previousUser };
      },
      onSuccess: (response) => {
        queryClient.setQueryData(["user"], () => [response.data]);
      },
      onError: (err, _ , context) => {
        console.log(err);
        queryClient.setQueryData(["user"], context?.previousUser);
      }
    });
  };

  return [getUser, setUser] as const;
};

export default useUser;
export type { User };
