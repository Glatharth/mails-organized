import re

input_file = "input.txt"


def getEmails(file):
    file_input = open(file, "r")
    var = re.findall("[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,4}", file_input.read())
    file_input.close()
    return var


def getEmailsQuota(file):
    file_input = open(file, "r")
    var = re.findall(".*\/.*\/.*", file_input.read())
    file_input.close()
    return var


def organize(file):
    quotas = getEmailsQuota(file)
    list_emails = getEmails(file)
    emails = str()
    print(len(list_emails))
    for i in range(len(list_emails)):
        print(i)
        emails += list_emails[i] + ": " + quotas[i] + "\n"

    var = "Quantidade de emails: " + str(len(list_emails)) + "\n\n" + str(emails)

    return var


emails_organized = organize(input_file)

print(emails_organized)

file_output = open("output.txt", "w+")
file_output.write(emails_organized)
file_output.close()
