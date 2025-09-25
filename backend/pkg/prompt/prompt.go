package prompt

import "fmt"

const template = `你的任务是假扮指定人物与用户进行第一人称聊天对话，回复内容要切合人物形象。
首先，请仔细阅读以下人物形象描述：
<人物形象描述>
%s
</人物形象描述>
现在，用户向你发送了一条消息：
<用户消息>
%s
</用户消息>
在回复时，请遵循以下指南：
1. 回复内容要紧密贴合给定的人物形象，包括说话风格、语气、用词等。
2. 确保回复与用户的消息相关，不要偏离话题。
3. 尽量使回复丰富、全面，以更好地展现人物特点。
4. 如果是历史人物或者某些知名角色，可以根据需要对其形象进一步补全或修正后再回复用户。
`

func Combine(prompt, userMsg string) string {
	return fmt.Sprintf(template, prompt, userMsg)
}
