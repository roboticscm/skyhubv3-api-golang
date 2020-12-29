package locale_resource

import (
	. "backend/system/error"
	. "backend/system/models"
	"bytes"
	"log"
	"net/http"
	"strings"
	"text/template"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
)

func saveHandler(c echo.Context) error {
	o := orm.NewOrm()
	lang := &LocaleResource{}
	if err := c.Bind(lang); err != nil {
		return BindObjectError(c, err, "LANG")
	}
	newId, err := o.Insert(lang)
	if err != nil {
		return SaveObjectError(c, err, "LANG")
	}

	savedObj := &LocaleResource{Id: newId}
	if err := o.Read(savedObj); err != nil {
		return LoadObjectError(c, err, "LANG")
	}

	return c.JSON(http.StatusOK, savedObj)
}

func updateHandler(c echo.Context) error {
	o := orm.NewOrm()
	newLang := &LocaleResource{}
	if err := c.Bind(newLang); err != nil {
		return BindObjectError(c, err, "LANG")
	}
	oldLang := &LocaleResource{Id: newLang.Id}
	if err := o.Read(oldLang); err != nil {
		return LoadObjectError(c, err, "LANG")
	}

	oldLang.Value = newLang.Value
	if _, err := o.Update(oldLang); err != nil {
		return UpdateObjectError(c, err, "LANG")
	}

	return c.JSON(http.StatusOK, oldLang)
}

func deleteHandler(c echo.Context) error {
	o := orm.NewOrm()
	newLang := &LocaleResource{}
	if err := c.Bind(newLang); err != nil {
		return BindObjectError(c, err, "LANG")
	}
	oldLang := &LocaleResource{Id: newLang.Id}
	if err := o.Read(oldLang); err != nil {
		return LoadObjectError(c, err, "LANG")
	}

	willDeleted := *oldLang
	if _, err := o.Delete(oldLang); err != nil {
		return DeleteObjectError(c, err, "LANG")
	}

	return c.JSON(http.StatusOK, willDeleted)
}

func getHandler(c echo.Context) error {
	o := orm.NewOrm()
	langs := []LocaleResource{}

	if _, err := o.QueryTable("locale_resource").All(&langs); err != nil {
		return LoadObjectError(c, err, "LANG")
	}
	return c.JSON(http.StatusOK, langs)
}

func getInitialHandler(c echo.Context) error {
	locale := c.QueryParam("locale")
	if locale == "" {
		return QueryParamError(c, "locale query param is invalid or missing", "LANG")
	}
	o := orm.NewOrm()
	langs := []LocaleResource{}

	if _, err := o.Raw("select * from find_language(?, ?)", nil, locale).QueryRows(&langs); err != nil {
		return LoadObjectError(c, err, "LANG")
	}

	return c.JSON(http.StatusOK, langs)
}

func generateReportHanlder(c echo.Context) error {
	testTemplate, err := template.ParseFiles("features/locale_resource/report.html")
	if err != nil {
		panic(err)
	}
	o := orm.NewOrm()
	langs := []LocaleResource{}

	if _, err := o.QueryTable("locale_resource").All(&langs); err != nil {
		return LoadObjectError(c, err, "LANG")
	}

	w := c.Response().Writer
	err = testTemplate.Execute(w, langs)

	return err
}

func generatePdfHanlder(c echo.Context) error {
	testTemplate, err := template.ParseFiles("features/locale_resource/report.html")
	if err != nil {
		panic(err)
	}
	o := orm.NewOrm()
	langs := []LocaleResource{}

	if _, err := o.QueryTable("locale_resource").All(&langs); err != nil {
		return LoadObjectError(c, err, "LANG")
	}

	buffer := &bytes.Buffer{}
	err = testTemplate.Execute(buffer, langs)
	if err != nil {
		panic(err)
	}

	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return err
	}
	page := wkhtml.NewPageReader(strings.NewReader(strings.Replace(buffer.String(), `<link rel="stylesheet" type="text/css" href="/css/report-preview.css">`, "", -1)))
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(20)
	page.Zoom.Set(0.95)
	page.HeaderCenter.Set("Header")

	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	//Your Pdf Name
	err = pdfg.WriteFile("Your_pdfname.pdf")
	if err != nil {
		log.Fatal(err)
	}

	r := c.Request()
	w := c.Response().Writer
	// pdfg.SetOutput(w)
	w.Header().Set("Content-Disposition", "attachment; filename=WHATEVER_YOU_WANT.pdf")
	w.Header().Set("Content-Type", "application/pdf")

	http.ServeFile(w, r, "Your_pdfname.pdf")

	return err
}
