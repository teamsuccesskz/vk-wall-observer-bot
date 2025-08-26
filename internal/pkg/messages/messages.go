package messages

const StartMessage = "👋 Hi, *%s!*\n\n" +
	"I keep track of new posts in VK communities and profiles.\n\n" +
    "📌 First add the IDs of the communities or profiles you are interested in.\n\n" +
    "⏰ Once an hour I will check if there are any new posts.\n\n" +
    "📩 If there are any, I will send you messages with their contents.\n\n"
const MenuMessage = "👨‍💻 *Choose an action...*\n\n To add/remove a community or profile, enter its URL or ID.\n\n*Example:* `https://vk.com/vk ` or `vk`"
const AddSlugSuccessful = "✅ ID `%s` added!"
const DeleteSlugSuccessful = "✅ Identifier `%s` has been deleted!"
const SlugNotFound = "❌ The ID of `%s` was not found!"
const SlugAlreadyExists = "❌ The ID of `%s` already exists!"
const SlugsListIsEmpty ="ℹ️ The list of IDs is empty!"
const SlugsListTitle ="📝 *List of IDs:*\n\n"
const SlugsListMessage = "%d. [%s](https://vk.com/%s) /// `%s`\n\n"
const AddSlugTooltip = "Enter the URL or ID to start tracking the community or profile"
const DeleteSlugTooltip = "Enter the URL or ID to delete the community or profile"
