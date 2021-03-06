package repos

import (
	"strings"

	"github.com/cbroglie/mustache"
	"github.com/khades/servbot/models"
	"gopkg.in/mgo.v2/bson"
)

var templateCollection = "templates"

func SetChannelTemplateAlias(user *string, userID *string, channelID *string, commandName *string, aliasTo *string) {
	aliasStripped := strings.ToLower(strings.Join(strings.Fields(*aliasTo), ""))
	commandNameStripped := strings.ToLower(strings.Join(strings.Fields(*commandName), ""))
	result, error := GetChannelTemplate(channelID, &aliasStripped)
	aliasTemplate := models.TemplateInfoBody{}
	if error == nil {
		aliasTemplate = result.TemplateInfoBody
	}

	putChannelTemplate(user, userID, channelID, &commandNameStripped, &aliasTemplate)
	PushCommandsForChannel(channelID)

}
func SetChannelTemplate(user *string, userID *string, channelID *string, commandName *string, template *models.TemplateInfoBody) error {
	commandNameStripped := strings.ToLower(strings.Join(strings.Fields(*commandName), ""))
	if template.Template == "" {
		_, templateError := mustache.ParseString(template.Template)
		if templateError != nil {
			return templateError
		}
	}
	putChannelTemplate(user, userID, channelID, &commandNameStripped, template)
	PushCommandsForChannel(channelID)
	return nil
}

func GetChannelTemplate(channelID *string, commandName *string) (models.TemplateInfo, error) {
	var result models.TemplateInfo
	error := Db.C(templateCollection).Find(models.TemplateSelector{ChannelID: *channelID, CommandName: *commandName}).One(&result)
	return result, error
}

// GetChannelTemplateWithHistory gets specific paginated
func GetChannelTemplateWithHistory(channelID *string, commandName *string) (*models.TemplateInfoWithHistory, error) {
	var result models.TemplateInfoWithHistory
	error := Db.C(templateCollection).Find(models.TemplateSelector{ChannelID: *channelID, CommandName: *commandName}).One(&result)
	return &result, error
}

func GetChannelTemplates(channelID *string) (*[]models.TemplateInfo, error) {
	var result []models.TemplateInfo
	error := Db.C(templateCollection).Find(models.ChannelSelector{ChannelID: *channelID}).All(&result)
	return &result, error
}

func GetChannelActiveTemplates(channelID *string) (*[]models.TemplateInfo, error) {
	var result []models.TemplateInfo
	error := Db.C(templateCollection).Find(bson.M{"channelid": *channelID, "template": bson.M{"$ne": ""}}).All(&result)
	return &result, error
}

func GetChannelAliasedTemplates(channelID *string, aliasTo *string) ([]models.TemplateInfo, error) {
	var result []models.TemplateInfo
	error := Db.C(templateCollection).Find(models.TemplateAliasSelector{ChannelID: *channelID, AliasTo: *aliasTo}).All(&result)
	return result, error
}
