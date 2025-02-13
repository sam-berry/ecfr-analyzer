import { NotificationData, notifications } from "@mantine/notifications";

function errorNotification({
  title,
  message,
  keepOpen,
}: {
  title: string;
  message?: string;
  keepOpen?: boolean;
}) {
  const props = {
    title,
    message,
    color: "red",
  } as NotificationData;

  if (keepOpen) {
    props.autoClose = false;
  }

  notifications.show(props);
}

function successNotification({
  title,
  message,
}: {
  title: string;
  message?: string;
}) {
  notifications.show({
    title,
    message,
    color: "green",
  } as NotificationData);
}

export { errorNotification, successNotification };
