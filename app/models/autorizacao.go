package models

import (
	"encoding/json"
	"fmt"

	"github.com/nathanzeras/go-ws-tks/config"
)

type Autorizacao struct {
	NumeroguiaPrestador        *string `json:"numeroGuiaPrestador"`
	CodigoAutorizacao          int    `json:"codigoAutorizacao"`
	SenhaAutorizacao           *string `json:"senhaAutorizacao"`
	DataAutorizacao            *string `json:"dataAutorizacao"`
	DescricaoGlosa             *string `json:"descricaoGlosa"`
	CodigoGlosa                *string `json:"codigoGlosa"`
	DataExpiracaoAutorizacao   *string `json:"dataExpiracaoAutorizacao"`
	TipoGuia                   *string `json:"tipoGuia"`
	CodigoProcedimento         int    `json:"codigoProcedimento"`
	DescricaoProcedimento      *string `json:"descricaoProcedimento"`
	QuantidadeSolicitada       int    `json:"quantidadeSolicitada"`
	QuantidadeAutorizada       int    `json:"quantidadeAutorizada"`
	QuantidadeRealizada        int    `json:"quantidadeRealizada"`
	DataExecucao               *string `json:"dataExecucao"`
	CodigoStatus               *string `json:"codigoStatus"`
	NumeroConselhoProfissional *string `json:"numeroConselhoProfissional"`
	ConselhoProfissional       *string `json:"conselhoProfissional"`
	NomeProfissional           *string `json:"nomeProfissional"`
	NomeBeneficiario           *string `json:"nomeBeneficiario"`
	Cpf                        *string `json:"cpf"`
	NumeroCarteira             *string `json:"numeroCarteira"`
	NumeroGuiaPrincipal        *string `json:"numeroGuiaPrincipal"`
	ValidadeCarteira           *string `json:"validadeCarteira"`
	Cbos                       *string `json:"cbos"`
	DataSolicitacao            *string `json:"dataSolicitacao"`
	CodigoContratadoSolicitante *string `json:"codigoContratadoSolicitante"`
	NomeContratadoSolicitante *string `json:"nomeContratadoSolicitante"`
	ValidadeSenha             *string `json:"validadeSenha"`
	NumeroGuiaOperadora       *string `json:"numeroGuiaOperadora"`
}

type Envio struct {
	NomeBeneficiario   string  `json:"nomeBeneficiario"`
	Cpf                *string `json:"cpf"`
	DataNascimento     string  `json:"dataNascimento"`
	CodigoProcedimento string  `json:"codigoProcedimento"`
	DataAtendimento    string  `json:"dataAtendimento"`
}

type RetornoAutorizacao struct {
	Status    bool   `json:"Status"`
	Descricao string `json:"Descricao"`
}

func SearchAutorizacao(e *Envio) []byte {

	db := config.ConnectDB()
	defer db.Close()
	var retornoAutorizacao RetornoAutorizacao

	//i := 0
	rows, err := db.Query(`select
	pe.ds_guia_prestador as numeroGuiaPrestador,
	pe.cd_autorizacao as codigoAutorizacao,
	case
		pe.ds_guia_principal when '' then pe.ds_guia_prestador
		else pe.ds_guia_principal
	end as senhaAutorizacao,
	case
		ex.dt_procedimento when '' then pe.dt_autorizacao
		else ex.dt_procedimento
	end as dataAutorizacao,
	pe.ds_glosa as descricaoGlosa,
	pe.cd_glosa as codigoGlosa,
	pe.dt_expiracao_autorizacao as dataExpiracaoAutorizacao,
	ex.cd_procedimento as codigoProcedimento,
	ex.ds_procedimento as descricaoProcedimento,
	ex.qtd_solicitada as quantidadeSolicitada,
	ex.qtd_autorizada as quantidadeAutorizada,
	ex.qtd_realizada as quantidadeRealizada,
	ex.cd_status as codigoStatus,
	ex.dt_procedimento as dataExecucao,
	pe.ds_tipo_guia as tipoGuia,
	pe.ds_crm as numeroConselhoProfissional,
	pe.ds_conselho as conselhoProfissional,
	pe.ds_medico as nomeProfissional,
	pe.ds_beneficiario as nomeBeneficiario,
	pe.ds_cpf as cpf,
	pe.nr_carteira as numeroCarteira,
	case
		pe.ds_guia_principal when '' then pe.ds_guia_prestador
		else pe.ds_guia_principal
	end as numeroGuiaOperadora,
	pe.ds_guia_principal as numeroGuiaPrincipal,
	case
		pe.dt_validade_carteira when '' then null
		else pe.dt_validade_carteira
	end as validadeCarteira,
	pe.ds_cbos as cbos,
	case
		pe.dt_solicitacao_guia when '' then null
		else (pe.dt_solicitacao_guia::date)::varchar
	end as dataSolicitacao,
	pe.cd_solicitante as codigoContratadoSolicitante,
	pe.ds_solicitante as nomeContratadoSolicitante,
	case
		pe.dt_validade_senha when '' then null
		else pe.dt_validade_senha
	end as validadeSenha
from pedidos pe
	  join exames ex on pe.ds_guia_prestador = ex.ds_guia_prestador
where
 (pe.ds_cpf = $1 or 
 ((SIMILARITY(pe.ds_beneficiario ,$4) > 0.5) and 
 (case
	when pe.dt_nascimento = '' then null
	else pe.dt_nascimento 
 end)::date = to_date($5, 'YYYYMMDD') )) and 
 ex.cd_procedimento = $2
 and ((pe.created_at)::date) >= to_date($3, 'YYYYMMDD') 
 order by ex.created_at desc limit 1`, e.Cpf, e.CodigoProcedimento, e.DataAtendimento, e.NomeBeneficiario, e.DataNascimento)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	//fmt.Println(CdPaciente)

	autorizacao := Autorizacao{}

	for rows.Next() {
		if err := rows.Scan(&autorizacao.NumeroguiaPrestador, &autorizacao.CodigoAutorizacao, 
			&autorizacao.SenhaAutorizacao, &autorizacao.DataAutorizacao, &autorizacao.DescricaoGlosa, 
			&autorizacao.CodigoGlosa, &autorizacao.DataExpiracaoAutorizacao, &autorizacao.CodigoProcedimento, 
			&autorizacao.DescricaoProcedimento, &autorizacao.QuantidadeSolicitada, &autorizacao.QuantidadeAutorizada,
			&autorizacao.QuantidadeRealizada, &autorizacao.CodigoStatus, &autorizacao.DataExecucao,
			&autorizacao.TipoGuia, &autorizacao.NumeroConselhoProfissional, &autorizacao.ConselhoProfissional,
			&autorizacao.NomeProfissional, &autorizacao.NomeBeneficiario, &autorizacao.Cpf, &autorizacao.NumeroCarteira,&autorizacao.NumeroGuiaOperadora,
			&autorizacao.NumeroGuiaPrincipal, &autorizacao.ValidadeCarteira, &autorizacao.Cbos, &autorizacao.DataSolicitacao,
			&autorizacao.CodigoContratadoSolicitante, &autorizacao.NomeContratadoSolicitante, &autorizacao.ValidadeSenha); err != nil {
			fmt.Println(err)
		}
	}

	if autorizacao.QuantidadeSolicitada > 0 {
		output, err := json.Marshal(autorizacao)
		if err != nil {
			fmt.Println("Error marshalling to json:", err)
		}
		//fmt.Print(string(output))
		return output
	} else {
		retornoAutorizacao.Status = false
		retornoAutorizacao.Descricao = "Autorização não encontrada."

		output, err := json.Marshal(retornoAutorizacao)
		if err != nil {
			fmt.Println("Error marshalling to json:", err)
		}
		//fmt.Print(string(output))
		return output
	}

	/*if err != nil {
		fmt.Println(err)
		//return patient, i, err
	} else {
		for rows.Next() {
			err = rows.Scan(&autorizacao)
		}*/

	/*_, err = dbr.Load(rows, &autorizacao)

	output, err := json.Marshal(autorizacao)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
	//}*/
	//return output
}

func SearchAutorizacaoOld(p map[string]string) []byte {

	
	db := config.ConnectDB()
	defer db.Close()
	var retornoAutorizacao RetornoAutorizacao

	//i := 0
	rows, err := db.Query(`select
	pe.ds_guia_prestador        as numeroGuiaPrestador,
	pe.cd_autorizacao           as codigoAutorizacao,
	pe.ds_senha_autorizacao     as senhaAutorizacao,
	pe.dt_autorizacao           as dataAutorizacao,
	pe.ds_glosa                 as descricaoGlosa,
	pe.cd_glosa                 as codigoGlosa,
	pe.dt_expiracao_autorizacao as dataExpiracaoAutorizacao,
	ex.cd_procedimento          as codigoProcedimento,
	ex.ds_procedimento          as descricaoProcedimento,
	ex.qtd_solicitada           as quantidadeSolicitada,
	ex.qtd_autorizada           as quantidadeAutorizada,
	ex.cd_status                as codigoStatus,
	pe.ds_tipo_guia             as tipoGuia,
	pe.ds_crm                   as numeroConselhoProfissional,
	pe.ds_medico                as nomeProfissional,
	pe.ds_beneficiario          as nomeBeneficiario,
	pe.ds_cpf                   as cpf,
	pe.nr_carteira              as numeroCarteira
from pedidos pe
	  join exames ex on pe.ds_guia_prestador = ex.ds_guia_prestador
where
 pe.ds_cpf = $1 and ex.cd_procedimento = $2
 and ((pe.created_at)::date) >= to_date($3, 'YYYYMMDD')
 order by pe.created_at desc limit 1`, p["cpf"], p["procedimento"], p["dt_data"])

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	//fmt.Println(CdPaciente)

	autorizacao := Autorizacao{}

	for rows.Next() {
		if err := rows.Scan(&autorizacao.NumeroguiaPrestador, &autorizacao.CodigoAutorizacao, &autorizacao.SenhaAutorizacao,
			&autorizacao.DataAutorizacao, &autorizacao.DescricaoGlosa, &autorizacao.CodigoGlosa, &autorizacao.DataExpiracaoAutorizacao,
			&autorizacao.CodigoProcedimento, &autorizacao.DescricaoProcedimento, &autorizacao.QuantidadeSolicitada, &autorizacao.QuantidadeAutorizada,
			&autorizacao.CodigoStatus, &autorizacao.TipoGuia, &autorizacao.NumeroConselhoProfissional, &autorizacao.NomeProfissional,
			&autorizacao.NomeBeneficiario, &autorizacao.Cpf, &autorizacao.NumeroCarteira); err != nil {
			fmt.Println(err)
		}
	}

	if autorizacao.QuantidadeSolicitada > 0 {
		output, err := json.Marshal(autorizacao)
		if err != nil {
			fmt.Println("Error marshalling to json:", err)
		}
		return output
	} else {
		retornoAutorizacao.Status = false
		retornoAutorizacao.Descricao = "Autorização não encontrada."

		output, err := json.Marshal(retornoAutorizacao)
		if err != nil {
			fmt.Println("Error marshalling to json:", err)
		}
		return output
	}

	/*if err != nil {
		fmt.Println(err)
		//return patient, i, err
	} else {
		for rows.Next() {
			err = rows.Scan(&autorizacao)
		}*/

	/*_, err = dbr.Load(rows, &autorizacao)

	output, err := json.Marshal(autorizacao)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
	//}*/
	//return output
}
